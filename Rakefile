# Compile the code, passing an argument to the task will set the file output.
# the ARCH environemnt variable to force the arch. currently 
# x86-64 (default), 386 and arm are supported.
#
desc "Compile the program, set ARCH to force the architecture."
task :compile do |task, args|
	files = %W{chunked_writer.go http_client.go route.go main.go}

	# default architecture is x86-64
	arch = ENV["ARCH"] || "x86-64"
	output = (args.first || File.basename(File.dirname(__FILE__))) + "-#{arch}"
	# Compile
  err = `#{compiler(arch)} -o #{output}.6 #{files.join(' ')}`
  raise "Compilation failed: #{err}" unless err == ""
	# Link
	err = `#{linker(arch)} -o #{output}.out #{output}.6`
	raise "Linking failed: #{err}" unless err == ""

  puts "compiled to #{output}.out"
end

desc "Remove intermediate files (*.6, *.8, *.5)"
task :cleanup do
	require 'fileutils'
	Dir.glob("*.{5,6,8}").each do |f|
		puts "Deleting #{f}"
		FileUtils.rm(f)
	end
end


private

def compiler(arch)
	compiler_code(arch) + "g"
end

def linker(arch)
	compiler_code(arch) + "l"
end

def compiler_code(arch)
	case arch
	when /x86-64/i
		"6"
	when /386/
		"8"
	when /arm/i
		"5"
	else
		raise "Architecture (#{arch}) not currently supported."
	end
end
